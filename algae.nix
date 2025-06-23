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

  doCheck = false;

  preBuild = ''
  ${templ}/bin/templ generate
  '';
}
