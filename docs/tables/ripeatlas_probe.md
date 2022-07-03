# Table: ripeatlas_probe

A probe is a software hosted from volunteers all over the world that actively measure Internet connectivity through ping, traceroute, DNS, SSL/TLS, NTP and HTTP measurements.

## Examples

### Get one probe

```sql
select
  address_v4,
  is_public
from
  ripeatlas_probe
where
  id=1
```
