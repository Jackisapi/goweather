# shell.nix
{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  name = "go-dev-shell";

  buildInputs = [
    pkgs.go
    pkgs.gopls     # optional: Go language server for IDEs
  ];

  shellHook = ''
    export GOPATH=$(pwd)/.gopath
    export GOBIN=$GOPATH/bin
    export PATH=$GOBIN:$PATH


    echo "Go dev environment ready"
  '';
}
