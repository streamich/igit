#!/usr/bin/env node

const { spawn } = require('child_process');
const { join } = require('path');
const { platform } = require('os');

const os = platform();
const cwd = process.cwd();
const filename = (os === 'win32' ? 'windows' : os) + '_amd64' + (os === 'win32' ? '.exe' : '');
const command = join(cwd, 'dist', filename);
const [,, ...args] = process.argv;

spawn(command, args, {
  cwd,
  stdio: 'inherit',
});
