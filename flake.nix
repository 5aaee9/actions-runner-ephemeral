{
  description = "A very basic flake";

  inputs = {
    indexyz.url = "github:X01A/nixos";
    nixpkgs.follows = "indexyz/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";

    devshell = {
      url = "github:numtide/devshell";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "aarch64-linux"
        "x86_64-linux"
      ];

      imports = [
        inputs.devshell.flakeModule
        inputs.treefmt-nix.flakeModule

        ./flake/dev.nix
        ./flake/fmt.nix
      ];
    };
}