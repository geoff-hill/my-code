# yow2021-k8s-native-java
A bunch of demos inspired by Josh Long's excellent presentation 'Kubernetes Native Java' at the YOW 2021 online conference.
Josh live-coded a lot of this in the presentation, and never stopped explaining even for a second.

So I basically typed in a lot of the code he was creating as I re-listened to his presentation, as a learning exercise.
I have however made a few changes, so that is the value of this repo. 
If anything looks good, it was probably Josh that created it, and if it doesn't work, it is probably my mistake.

The main changes are:
- Used gradle instead of maven to build, as I'm over maven;
- Made some different tooling decisions around how to build containers;
- Created a mono-repo for the whole thing, built as a gradle multi-module project;
- Added in some principles around K8S deployment that I've learned in other places.

Hope it works for you, enjoy.
