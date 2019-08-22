with import <nixpkgs> {};

let
  libkirk = (import (fetchTarball https://github.com/juliosueiras/nix-vita/archive/master.tar.gz) {}).vitaPackages.libkirk;
in buildGoPackage {
  name = "chovy-sign-cli";

  src = ./.;

  goPackagePath = "github.com/juliosueiras/chovy-sign-cli";

  CGO_LDFLAGS = "-lkirk";

  nativeBuildInputs = [
    autoPatchelfHook
  ];

  buildInputs = [
    libkirk
  ];
  
  preInstall = if stdenv.isDarwin then ''
    install_name_tool -change libkirk.so ${libkirk}/lib/libkirk.so go/bin/chovy-sign-cli
  '' else "";

  doCheck = false;
}

