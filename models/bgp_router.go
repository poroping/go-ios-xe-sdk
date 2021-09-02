package models

type BgpRouter struct {
	Bgp struct {
		ID  string `json:"id,omitempty"`
		Bgp struct {
			AigpRibMetric         string `json:"aigp-rib-metric,omitempty"`
			AlwaysCompareMed      string `json:"always-compare-med,omitempty"`
			ClusterID             string `json:"cluster-id,omitempty"`
			DeterministicMed      string `json:"deterministic-med,omitempty"`
			EnforceFirstAsBoolean bool   `json:"enforce-first-as-boolean,omitempty"`
			EnforceFirstAs        string `json:"enforce-first-as,omitempty"`
			EnhancedError         bool   `json:"enhanced-error,omitempty"`
			FastExternalFallover  bool   `json:"fast-external-fallover,omitempty"`
			LogNeighborChanges    bool   `json:"log-neighbor-changes,omitempty"`
			MaxasLimit            int    `json:"maxas-limit,omitempty"`
			MaxcommunityLimit     int    `json:"maxcommunity-limit,omitempty"`
			RouteMapCache         bool   `json:"route-map-cache,omitempty"`
			SafeEbgpPolicy        string `json:"safe-ebgp-policy,omitempty"`
			UpdateDelay           int    `json:"update-delay,omitempty"`
			KeepaliveInterval     int    `json:"keepalive-interval,omitempty"`
			Holdtime              int    `json:"holdtime,omitempty"`
			MinimumNeighborHold   int    `json:"minimum-neighbor-hold,omitempty"`
		} `json:"bgp,omitempty"`
	} `json:"Cisco-IOS-XE-bgp:bgp,omitempty"`
}

type BgpRouterList struct {
	Bgp []struct {
		ID  string `json:"id,omitempty"`
		Bgp struct {
			AigpRibMetric         string `json:"aigp-rib-metric,omitempty"`
			AlwaysCompareMed      string `json:"always-compare-med,omitempty"`
			ClusterID             string `json:"cluster-id,omitempty"`
			DeterministicMed      string `json:"deterministic-med,omitempty"`
			EnforceFirstAsBoolean bool   `json:"enforce-first-as-boolean,omitempty"`
			EnforceFirstAs        string `json:"enforce-first-as,omitempty"`
			EnhancedError         bool   `json:"enhanced-error,omitempty"`
			FastExternalFallover  bool   `json:"fast-external-fallover,omitempty"`
			LogNeighborChanges    bool   `json:"log-neighbor-changes,omitempty"`
			MaxasLimit            int    `json:"maxas-limit,omitempty"`
			MaxcommunityLimit     int    `json:"maxcommunity-limit,omitempty"`
			RouteMapCache         bool   `json:"route-map-cache,omitempty"`
			SafeEbgpPolicy        string `json:"safe-ebgp-policy,omitempty"`
			UpdateDelay           int    `json:"update-delay,omitempty"`
			KeepaliveInterval     int    `json:"keepalive-interval,omitempty"`
			Holdtime              int    `json:"holdtime,omitempty"`
			MinimumNeighborHold   int    `json:"minimum-neighbor-hold,omitempty"`
		} `json:"bgp,omitempty"`
	} `json:"Cisco-IOS-XE-bgp:bgp,omitempty"`
}