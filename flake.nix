{
  description = "Devshell for the Mistral Go SDK";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAll = nixpkgs.lib.genAttrs systems;
    in
    {
      devShells = forAll (system:
        let
          # speakeasy-cli is marked unfree in nixpkgs
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfreePredicate = pkg: pkg.pname or "" == "speakeasy-cli";
          };
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              speakeasy-cli
              openapi-generator-cli
              oapi-codegen
              gnused
              gh
              gopls
            ];
          };
        });
    };
}
