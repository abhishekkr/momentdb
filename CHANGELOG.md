
v0.0.1

* it's like vanilla GoShare

---

### Good To Have in Changelog

* get a proc for MitM(~Splitter) of packets to default GoShare
* make ~Splitter check which GoShare to send, slowly add more ways to shard-logic; have ~Split/~Join develop together
* make ~NodeCop check for all marked nodes at regular interval and turn them Up/Down
* make ~DataCop try for all Failed-Push for configured-times in-increasing-delay-retries; data stays in local-data-cop-shard; no try by ~DataCop if ~Node marked ~Down and re-queued for next interval
* make ~DataCop capable of copying over one ~Node onto another
* make ~Splitter apply ~Replicas; with Push/Delete to all (any Failure delegated to ~DataCop) and Read in load-balanced-mode
* make ~Splitter delegate all Packets to MitM(~Sideshow) passing only ~SideAct(blank-to-start) it to ~Splitter.side
* make ~Splitter.side pass all Packets to ~Sideshow
* make ~Sideshow have some basic ~SideAct like ~ChildFactor from existing ~Factor(s); like "{eq|ne|ge|gt|le|lt}-a-value" branches [awesome to have Failure|Warning|Best|Worst|* channels]
* make ~Sideshow to (like map-reduce) apply SOME_FUNC over data-stream(s); in Bucket-Quantity of data-pool(s); within certain Time-Range
* make ~Sideshow accept Custom ~SideAct

* a defer-cleanup for series task, picked up when have enough free resource

* support for grafana

* providing self-stat

* providing 'join' capability across series

* providing time-range capability

* [improving speed of mathematical actions over time-range series] can turn-on average,total,etc on series per-minute; calculating to per-hour; calculating to per-day; calculating to per-month; calculating to per-year



---
