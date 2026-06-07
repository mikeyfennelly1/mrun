{
  description = "mrun - OCI-compatible Linux container runtime";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      # mrun uses Linux-only syscalls (namespaces, cgroups, capabilities).
      systems = [ "x86_64-linux" "aarch64-linux" ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in {
      packages = forEachSystem (pkgs: {
        default = pkgs.buildGoModule {
          pname = "mrun";
          version = "0.1.0";
          src = ./.;

          # Run `nix build` once with this placeholder; replace with the hash
          # printed in the error: "got: sha256-<hash>".
          vendorHash = pkgs.lib.fakeHash;

          # mrun requires file capabilities (CAP_SYS_ADMIN etc.) set via
          # setcap after install. The Nix sandbox cannot run setcap, so do it
          # as a post-install step in your NixOS configuration or manually:
          #   sudo setcap <caps> $(which mrun)
          meta = with pkgs.lib; {
            description = "OCI-compatible Linux container runtime";
            license = licenses.mit;
            platforms = platforms.linux;
            mainProgram = "mrun";
          };
        };
      });

      devShells = forEachSystem (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go_1_25      # go.mod requires >=1.23; 1.25 is backward-compatible
            go-task      # Taskfile runner
            bats         # bash test framework (tests/build.bats)
            libcap       # provides setcap / getcap for the build script
            pkg-config   # needed by some cgo transitive deps
            git
          ];

          shellHook = ''
            export GOPATH="$HOME/go"
            export PATH="$PATH:$GOPATH/bin"
          '';
        };
      });
    };
}
