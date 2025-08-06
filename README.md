# TAR Integration

## What is TAR?

**TAR (Tape Archive)** is a widely used file format and utility for collecting many files into a single archive file, often used for backup and distribution.

This integration allows:

- Seamless backup of files and directories as TAR archives into a Plakar repository
- Direct restoration of snapshots as TAR archives to specified destinations
- Compatibility with systems and tools that use the TAR format

## Installation

**This integration is included in the default Plakar installation. No additional steps are required to enable it.**

## Configuration

The configuration parameters are as follow:

- `file` (required): Path to the TAR archive file

## Example Usage

```bash
# configure a TAR source
$ plakar source add myTARSrc tar:///path/to/archive.tar

# backup the source
$ plakar at /tmp/store backup @myTARSrc
```
