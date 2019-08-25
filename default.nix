with import <nixpkgs> {};

let
  vitaPackages = (import (fetchTarball https://github.com/juliosueiras/nix-vita/archive/58cee720943d81e91a9d62d068d883b25789d48e.tar.gz) {}).vitaPackages;
  projectPath = "github.com/juliosueiras/chovy-sign-cli";
in buildGoPackage {
  name = "chovy-sign-cli";

  src = ./.;

  goPackagePath = projectPath;

  CGO_LDFLAGS = "-lkirk";

  nativeBuildInputs = [
    autoPatchelfHook
  ];

  buildInputs = [
    vitaPackages.libkirk
  ];

  sign_np = "${vitaPackages.signnp}/bin/sign_np";

  preBuild = ''
    substituteAllInPlace go/src/${projectPath}/pbp/main.go
  '';

  
  preInstall = if stdenv.isDarwin then ''
    install_name_tool -change libkirk.so ${vitaPackages.libkirk}/lib/libkirk.so go/bin/chovy-sign-cli
  '' else "";

  doCheck = false;
}

