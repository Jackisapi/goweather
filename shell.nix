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
    export WEATHER_KEY=25ef63d3735611db59c5f7d82e6a814f


    echo "Go dev environment ready"
  '';
}
