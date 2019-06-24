{ pkgs ? import <nixpkgs> {}, ... }:
let
  inherit (pkgs) buildGoModule lib;
in

buildGoModule {
  name = "polybar-modules";
  src = ./.;
  modSha256 = "0nhfdjk642r58z4kpq8mq4ci14pavwlw7nrfp5xvybjj08kgsgfm";
}
