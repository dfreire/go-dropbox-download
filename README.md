Download files from a shared dropbox folder link.

# Install

## Option 1: instal the binary release

Download the release zip file for the system of your choice from the [releases page](https://github.com/dfreire/go-dropbox-download/releases/).

Unzip the file and place it somewhere on your system's PATH.

## Option 2: use your local go environment

```
$ go get github.com/dfreire/go-dropbox-download
$ cd <gopath>/src/github.com/dfreire/go-dropbox-download
$ go install go-dropbox-download
```

# Usage

```
$ go-dropbox-download <dropbox_folder_link> <local_folder> <match_filenames_string>
```

It will skip downloading files that are already present in the local folder.

# Example

```
$ go-dropbox-download "https://www.dropbox.com/sh/D3mb-t-BKH_aAUQ9/3FK3Txo_4f-rQbdb/MyPhotos?dl=0" "./MyPhotos" ".jpg"
```

