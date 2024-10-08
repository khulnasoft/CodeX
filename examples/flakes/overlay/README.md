# Adding Overlays with Flakes

For a more in depth walkthrough of this example, check out our [blog post](https://www.khulnasoft/blog/using-nix-flakes-with-codex/)

This flake shows how to use an overlay Nix flake to override a Nixpkgs package and use it in your codex configuration.

In this example, using the default `yarn` from Nixpkgs will cause `yarn start` to fail. To fix this issue, we use an overlay to modify the `yarn` package to use NodeJS 16 instead of it's default NodeJS 14.

```nix

   overlay = (final: prev: {
      yarn = prev.yarn.override { nodejs = final.pkgs.nodejs-16_x; };
   });
```

The yarn-overlay flake exports the modified yarn package in it's outputs. We can then use this package in our codex shell by adding it to `packages` in our `codex.json` file.

```json
{
   "packages": [
      "path:./yarn-overlay#yarn"
      "fnm"
   ]
   ...
}
```

Note: you will need Codex 0.4.7-dev or later for this to work. You can try it out by exporting `CODEX_VERSION=0.4.7-dev` before running `codex shell`.

You can use the flake.nix in the yarn-overlay directory as a template for creating your own overlays.
