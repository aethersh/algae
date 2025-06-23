{
  buildGoApplication,
  templ,
  pkgs, 
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
  ${pkgs.tailwindcss_4}/bin/tailwindcss -i templates/style.css -o static/style.css -m	
  '';
}
