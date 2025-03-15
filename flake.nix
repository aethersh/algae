{
  description = "AetherNet Looking Glass";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";

    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    { nixpkgs, rust-overlay, ... }:
    let

      supportedSystems = [
        "x86_64-linux"
        "aarch64-darwin"
      ];

      forEachSupportedSystem =
        f:
        nixpkgs.lib.genAttrs supportedSystems (
          system:
          f {
            pkgs = import nixpkgs {
              inherit system;
              overlays = [ (import rust-overlay) ];
            };
            inherit system;
          }
        );
    in

    {

      devShells = forEachSupportedSystem (
        { pkgs, ... }:
        {
          default = pkgs.mkShell rec {
            buildInputs = with pkgs; [
              openssl
              protobuf
              rust-bin.stable.latest.default
              rust-analyzer
            ];
            # LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath buildInputs;
          };
        }
      );
    };
}