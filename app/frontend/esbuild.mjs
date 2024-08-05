import esbuild from "esbuild";
import { spawn, spawnSync } from "child_process";
import fs from 'fs-extra';

const prodBuild = process.env.BUILD === 'true';
const buildDir = prodBuild ? 'dist' : 'build';

const runTailwindBuild = (watch) => {
  console.log("Building Tailwind CSS...");
  try {
    const command = 'npx';
    const args = [
      'tailwindcss',
      'build',
      '-i', 'src/styles/tailwind.css',
      '-o', `${buildDir}/site/tailwind.css`
    ];

    if (watch) {
      args.push('--watch');
      spawn(command, args, {
        stdio: 'inherit'
      });
    } else {
      spawnSync(command, args, {
        stdio: 'inherit'
      });
    }
    console.log("Tailwind CSS build successful!");
  } catch (error) {
    console.error("Error building Tailwind CSS:", error.message);
  }
};

const baseOptions = {
  bundle: true,
  loader: {
    ".ts": "tsx",
    ".tsx": "tsx",
    ".woff2": "file",
    ".woff": "file",
    ".html": "copy",
    ".json": "copy",
    ".ico": "copy",
  },
  minify: prodBuild,
  sourcemap: prodBuild ? false : "linked",
  define: {
    "global": "window",
    "process.env.BASE_URL": prodBuild ? '""' : '"http://45.55.132.24:8080"',
    "process.env.PRODUCTION": prodBuild ? '"true"' : '"false"'
  },
  entryNames: "[name]",
  logLevel: 'info',
};

async function doBuild(options) {
  runTailwindBuild(false);

  // Ensure the build directory exists
  fs.ensureDirSync(`${buildDir}/site/`);

  // Copy the static files
  fs.copySync('src/index.html', `${buildDir}/site/index.html`);

  await esbuild.build(options);
}

await doBuild({
  ...baseOptions,
  entryPoints: [
    "./src/index.tsx",
  ],
  outdir: `${buildDir}/site/`,
});
