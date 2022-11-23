package atomicasset

type SortOrder string

const (

	// SortNone order is not specified.
	SortNone SortOrder = ""

	// SortAscending sorts in ascending order
	SortAscending SortOrder = "asc"

	// SortDescending sorts in descending order.
	SortDescending SortOrder = "desc"
)

// LogRequestParams holds the parameters for an Log request
type LogRequestParams struct {
	Page            int       `qs:"page,omitempty"`
	Limit           int       `qs:"limit,omitempty"`
	Order           SortOrder `qs:"order,omitempty"`
	ActionWhitelist string    `qs:"action_whitelist,omitempty"`
	ActionBlacklist string    `qs:"action_blacklist,omitempty"`
}

// AssetsRequestParams holds the parameters for an Asset request
type AssetsRequestParams struct {
	CollectionName          string   `qs:"collection_name,omitempty"`
	CollectionBlacklist     []string `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist     []string `qs:"collection_whitelist,omitempty"`
	SchemaName              string   `qs:"schema_name,omitempty"`
	TemplateID              int      `qs:"template_id,omitempty"`
	TemplateWhitelist       []int    `qs:"template_whitelist,omitempty"`
	TemplateBlacklist       []int    `qs:"template_blacklist,omitempty"`
	Owner                   string   `qs:"owner,omitempty"`
	Match                   string   `qs:"match,omitempty"`
	MatchImmutableName      string   `qs:"match_immutable_name,omitempty"`
	MatchMutableName        string   `qs:"match_mutable_name,omitempty"`
	HideTemplatesByAccounts string   `qs:"hide_templates_by_accounts,omitempty"`

	IsTransferable          bool `qs:"is_transferable,omitempty"`
	IsBurnable              bool `qs:"is_burnable,omitempty"`
	Burned                  bool `qs:"burned,omitempty"`
	OnlyDuplicatedTemplates bool `qs:"only_duplicated_templates,omitempty"`
	HasBackedTokens         bool `qs:"has_backend_tokens,omitempty"`
	HideOffers              bool `qs:"hide_offers,omitempty"`

	LowerBound string `qs:"lower_bound,omitempty"`
	UpperBound string `qs:"upper_bound,omitempty"`

	Before int `qs:"before,omitempty"`
	After  int `qs:"after,omitempty"`

	Limit int    `qs:"limit,omitempty"`
	Order string `qs:"order,omitempty"`
	Sort  string `qs:"sort,omitempty"`
}

// AssetSalesRequestParams holds the parameters for an AssetSales request
type AssetSalesRequestParams struct {
	Buyer  string    `qs:"buyer,omitempty"`
	Seller string    `qs:"seller,omitempty"`
	Symbol string    `qs:"symbol,omitempty"`
	Order  SortOrder `qs:"order,omitempty"`
}

type CollectionSortColumn string

const (

	// CollectionSortDefault sorts by the default column (created)
	CollectionSortDefault CollectionSortColumn = ""

	// CollectionSortCreated sorts by the created column
	CollectionSortCreated CollectionSortColumn = "created"

	// CollectionSortName sorts by the collection name column
	CollectionSortName CollectionSortColumn = "collection_name"
)

type CollectionsRequestParams struct {
	Author            string               `qs:"author,omitempty"`
	Match             string               `qs:"match,omitempty"`
	AuthorizedAccount string               `qs:"authorized_account,omitempty"`
	NotifyAccount     string               `qs:"notify_account,omitempty"`
	Blacklist         string               `qs:"collection_blacklist,omitempty"`
	Whitelist         string               `qs:"collection_whitelist,omitempty"`
	IDs               string               `qs:"ids,omitempty"`
	LowerBound        string               `qs:"lower_bound,omitempty"`
	UpperBound        string               `qs:"upper_bound,omitempty"`
	Before            int                  `qs:"before,omitempty"`
	After             int                  `qs:"after,omitempty"`
	Page              int                  `qs:"page,omitempty"`
	Limit             int                  `qs:"limit,omitempty"`
	Order             SortOrder            `qs:"order,omitempty"`
	Sort              CollectionSortColumn `qs:"sort,omitempty"`
}
