{ pkgs ? import <nixpkgs> {}, ... }:
let
  inherit (pkgs) buildGoModule lib;
in

buildGoModule {
  name = "polybar-modules";
  src = ./.;
  modSha256 = "1faszp9hinm6ab13rvvc20cqbjxrlf3qpid0zs1frvf4w2y9fi8q";
}
