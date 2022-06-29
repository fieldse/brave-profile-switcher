## Brave Browser Profile Switcher

This is a simple GUI widget which allows for quickly opening a new Brave browser window 
in a selected profile.


### Rationale

A user may have multiple profiles on Brave browser (identical to Chrome/Chromium profiles).

Opening a new window in a specific profile is tedious, and there is no direct way to do so 
from the deskop. It requires *first* opening the default browser profile, then selecting
the profile from the menu, which opens up a *second* new window.

This GUI allows a one-hotkey switch to view all available profiles, and open a new window 
in one profile with one click.

### Requirements

**Run**
This is intended to run on the Linux desktop. It should run directly from a compiled binary,
and should not have any special requirements beyond an up-to-date kernel.

**Build**
- go version 1.18+


### Installation

This installs the binary to your /usr/local/bin path

```sh
$ ./install.sh
```

### Usage

Run from binary:

```sh
$ brave-profile-switcher

```

Run with go: 

```
$ go run cmd/brave-profile-switcher
```

### Testing

*TODO...*
