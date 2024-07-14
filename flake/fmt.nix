{ lib, inputs, ... }:
{
  imports = [ inputs.treefmt-nix.flakeModule ];

  perSystem =
    { pkgs, ... }:
    {
      treefmt = {
        # Used to find the project root
        projectRootFile = "flake.lock";
        programs.nixfmt = {
          enable = true;
          package = pkgs.nixfmt-rfc-style;
        };
      };
    };
}