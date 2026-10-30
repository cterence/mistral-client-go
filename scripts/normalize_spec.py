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
