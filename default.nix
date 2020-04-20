{ pkgs ? import <nixpkgs> {}, ... }:
let
  inherit (pkgs) buildGoModule lib;
in

buildGoModule {
  name = "polybar-modules";
  src = ./.;
  modSha256 = "1pzk6n245r7xz6lnjk41zg2n49aqs0l1a1xaw8fwbyxkmn9238dx";
}
