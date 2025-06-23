{
  description = "A declarative DNS server";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-25.05";

    flake-parts.url = "github:hercules-ci/flake-parts";

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    templ.url = "github:a-h/templ";
  };

  outputs = inputs @ {
    flake-parts,
    nixpkgs,
    gomod2nix,
    templ,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-darwin" "x86_64-darwin" "aarch64-linux"];
      perSystem = {
        pkgs,
        system,
        ...
      }: {
        # https://templ.guide/quick-start/installation#nix
        packages = rec {
          default = pkgs.callPackage ./algae.nix {
            inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
            inherit (inputs.templ.packages.${system}) templ;
          };
          algae = default;
        
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            air
            go
            gopls
            gotools
            go-tools
            gomod2nix.packages.${system}.default
            templ.packages.${system}.templ
            tailwindcss_4
	          just
          ];

          ALGAE_ALLOWED_ORIGINS = "http://localhost:2152";
          ALGAE_DOMAIN = "as215207.net";
          ALGAE_TEST_V6="2602:fbcf:df::1";
          ALGAE_LOCATION = "New York City, NY, USA";
        };
        formatter = pkgs.alejandra;
      };
    };
}
