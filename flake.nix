{
  description = "A cli tool to validate json against a schema";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-22.11";
  inputs.nixpkgs-unstable.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs =
    { self, nixpkgs, nixpkgs-unstable, flake-utils}:
    let
      pkgs = nixpkgs.legacyPackages;
      pkgs-unstable = nixpkgs-unstable.legacyPackages;
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
    in
    flake-utils.lib.eachDefaultSystem (system: {
      packages = {
        gojsonschema-cli = pkgs.${system}.buildGoModule {
          pname = "gojsonschema-cli";
          inherit version;
          src = ./.;
          vendorSha256 = "sha256-5IdyniX7aGxuZN9S8fsBAvDxuoLqG2jhfw6TsPXIEnI=";
        };
      };

      defaultPackage = self.packages.${system}.gojsonschema-cli;

      devShells.default = pkgs.${system}.mkShell {
        buildInputs = [
          pkgs-unstable.${system}.cobra-cli
          pkgs-unstable.${system}.go
        ];
      };
    });
}
