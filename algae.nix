{
  buildGoApplication,
  templ,
  ...
}:
buildGoApplication {
  pname = "algae";
  version = "0.0.0";

  src = ./.;
  pwd = ./.;
  modules = ./gomod2nix.toml;

  preBuild = ''
  ${templ}/bin/templ generate
  '';
}