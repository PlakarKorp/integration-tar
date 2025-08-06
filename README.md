# TAR Integration

## Overview

**TAR (Tape Archive)** is a widely used file format and utility for collecting many files into a single archive file, often used for backup and distribution.

This integration allows:

- Seamless backup of files and directories as TAR archives into a Plakar repository
- Direct restoration of snapshots as TAR archives to specified destinations
- Compatibility with systems and tools that use the TAR format

## Configuration

The configuration parameters are as follow:

- `location` (required): Path to the TAR archive file

> **Note:** The location can be write directly in the command, with `tar://` prefix. No need to add it in the configuration.

## Examples

```bash
# backup the source
$ plakar at /tmp/store backup tar:///tmp/example.tar
```
