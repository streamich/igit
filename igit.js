#!/usr/bin/env node

const { spawn } = require('child_process');
const { join } = require('path');
const { platform } = require('os');

const os = platform();
const filename = (os === 'win32' ? 'windows' : os) + '_amd64' + (os === 'win32' ? '.exe' : '');
const command = join(__dirname, 'dist', filename);
const [,, ...args] = process.argv;
const cwd = process.cwd();

spawn(command, args, {
  cwd,
  stdio: 'inherit',
});
