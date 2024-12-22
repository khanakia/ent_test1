```
{{ define "model/fields/additional" }}
    {{- if eq $.Name "Media" }}
        URL string `json:"url,omitempty"`
        MediableID string `json:"mediableID,omitempty"`
        Tag string `json:"tag,omitempty"`
        Order int `json:"order,omitempty"`
    {{- end }}

    {{- if eq $.Name "Post" }}
        Medias []*Media `json:"medias,omitempty"`
    {{- end }}
{{ end }}


{{ define "dialect/sql/query/fields/additional/url" }}
    withMedias      *MediaQuery
{{ end }}

{{ define "dialect/sql/query/all/nodes/withMedias" }}
    {{- $n := $ }}
    {{ if eq $n.Name "Post" }}
        if query := pq.withMedias; query != nil {
            if err := pq.loadMedias(ctx, query, nodes,
                func(n *Post) { n.Medias = []*Media{} },
                func(n *Post, e *Media) { n.Medias = append(n.Medias, e) }); err != nil {
                return nil, err
            }
        }
    {{ end }}


{{ end }}

// saasmedia.tmpl.go
func (uq *PostQuery) WithMedias(opts ...func(*MediaQuery)) *PostQuery {
	query := (&MediaClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withMedias = query
	return uq
}


func (uq *PostQuery) loadMedias(ctx context.Context, query *MediaQuery, nodes []*Post, init func(*Post), assign func(*Post, *Media)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Post)
	nids := make(map[string]map[*Post]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(mediable.Table)
		s.Join(joinT).On(s.C(media.FieldID), joinT.C(mediable.MediaColumn))
		s.Where(sql.InValues(joinT.C(mediable.FieldMediableID), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(mediable.FieldMediableID))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}

	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Post]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Media](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "workspaces" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
```