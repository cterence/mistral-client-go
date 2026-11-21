#!/usr/bin/env python3
"""Normalize the upstream OpenAPI spec for openapi-generator.

The published spec embeds JSON-Schema ``$defs`` blocks inside media-type
schemas with ``#/$defs/...`` references. openapi-generator's parser cannot
resolve those local references and fails spec validation. This script
hoists every ``$defs`` block into ``components.schemas`` (renaming on
collision) and rewrites the references to standard
``#/components/schemas/...`` form.
"""

import argparse
import sys

import yaml


def rewrite_refs(node, mapping):
    if isinstance(node, dict):
        ref = node.get("$ref")
        if isinstance(ref, str) and ref.startswith("#/$defs/"):
            name = ref[len("#/$defs/") :]
            if name not in mapping:
                raise ValueError(f"reference to unknown $defs entry: {ref}")
            node["$ref"] = f"#/components/schemas/{mapping[name]}"
        for key, value in list(node.items()):
            if (
                key != "$ref"
                and isinstance(value, str)
                and value.startswith("#/$defs/")
            ):
                name = value[len("#/$defs/") :]
                if name not in mapping:
                    raise ValueError(f"reference to unknown $defs entry: {value}")
                node[key] = f"#/components/schemas/{mapping[name]}"
            else:
                rewrite_refs(value, mapping)
    elif isinstance(node, list):
        for item in node:
            rewrite_refs(item, mapping)


def hoist_defs(node, mapping, components):
    if isinstance(node, dict):
        defs = node.pop("$defs", None)
        if defs:
            for name, schema in defs.items():
                target = name
                suffix = 0
                while target in components or target in mapping.values():
                    suffix += 1
                    target = f"{name}{suffix}"
                mapping[name] = target
                components[target] = schema
        for key in list(node):
            hoist_defs(node[key], mapping, components)
    elif isinstance(node, list):
        for item in node:
            hoist_defs(item, mapping, components)


def dedupe_tags(spec):
    # The generator emits one API group per tag; operations carrying several
    # tags produce duplicate request types that collide at compile time.
    # Keep only the most specific (last) tag, matching previous single-tag
    # behavior.
    for item in spec.get("paths", {}).values():
        for method, op in item.items():
            if method not in ("get", "put", "post", "delete", "options", "head", "patch", "trace"):
                continue
            tags = op.get("tags")
            if isinstance(tags, list) and len(tags) > 1:
                op["tags"] = [tags[-1]]


def int_bounds(node):
    # (low, high) of an integer schema node, or None. Handles OpenAPI 3.1
    # type arrays and composed schemas (anyOf/oneOf/allOf), where the
    # ``default`` sits on the composer node rather than the typed member.
    if not isinstance(node, dict):
        return None
    type_ = node.get("type")
    types = type_ if isinstance(type_, list) else [type_]
    format_ = node.get("format")
    is_int = any(
        isinstance(t, str) and (t.startswith("integer") or t in ("int32", "int64"))
        for t in types
    )
    if not is_int:
        for key in ("anyOf", "oneOf", "allOf"):
            for member in node.get(key) or []:
                bounds = int_bounds(member)
                if bounds is not None:
                    return bounds
        return None
    if "int32" in types or format_ == "int32":
        high = 2**31 - 1
    else:
        high = 2**63 - 1
    return (-high - 1, high)


def clamp_int_defaults(node):
    # Out-of-range integer defaults (e.g. 9223372036854776000 on an int32
    # field) become invalid Go literals in the generated code; clamp them to
    # the declared format's range. The default may sit on the schema, next
    # to an anyOf/oneOf/allOf composition, or on a parameter object next to
    # its schema.
    if isinstance(node, dict):
        bounds = int_bounds(node)
        if bounds is None and isinstance(node.get("schema"), dict):
            bounds = int_bounds(node["schema"])
        if bounds is None:
            # Width unknown; at least keep the literal representable in Go.
            bounds = (-(2**63 - 1) - 1, 2**63 - 1)
        low, high = bounds
        for key in ("default", "example"):
            value = node.get(key)
            if isinstance(value, int) and not isinstance(value, bool):
                node[key] = min(max(value, low), high)
        for value in node.values():
            clamp_int_defaults(value)
    elif isinstance(node, list):
        for item in node:
            clamp_int_defaults(item)


def rewrite_null_types(node):
    # Schemas declaring only ``type: null`` (dead fields) make the
    # generator emit an invalid ``*nil`` Go type. An empty schema (any
    # type) generates a valid NullableAnyOf equivalent.
    if isinstance(node, dict):
        if node.get("type") == "null":
            replacement = {k: v for k, v in node.items() if k != "type"}
            node.clear()
            node.update(replacement)
            return
        for value in node.values():
            rewrite_null_types(value)
    elif isinstance(node, list):
        for item in node:
            rewrite_null_types(item)


def normalize(spec):
    components = spec.setdefault("components", {}).setdefault("schemas", {})
    mapping = {}
    hoist_defs(spec, mapping, components)
    rewrite_refs(spec, mapping)
    rewrite_null_types(spec)
    clamp_int_defaults(spec)
    dedupe_tags(spec)
    return spec


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("spec", help="OpenAPI spec file to normalize in place")
    args = parser.parse_args()

    with open(args.spec) as f:
        spec = yaml.safe_load(f)

    normalize(spec)

    with open(args.spec, "w") as f:
        yaml.safe_dump(spec, f, sort_keys=False, allow_unicode=True, default_flow_style=False)
    return 0


if __name__ == "__main__":
    sys.exit(main())
