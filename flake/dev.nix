{
  perSystem =
    {
      inputs,
      config,
      pkgs,
      self,
      ...
    }:
    {
      devshells.default = {
        commands = [
          {
            package = pkgs.lefthook;
            category = "development";
          }
          {
            package = pkgs.nixfmt-rfc-style;
            category = "development";
          }
          {
            package = pkgs.go;
          }
        ];
      };
    };
}