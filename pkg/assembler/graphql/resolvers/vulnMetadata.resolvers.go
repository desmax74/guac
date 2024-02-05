package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/backends/helper"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestVulnerabilityMetadata is the resolver for the ingestVulnerabilityMetadata field.
func (r *mutationResolver) IngestVulnerabilityMetadata(ctx context.Context, vulnerability model.VulnerabilityInputSpec, vulnerabilityMetadata model.VulnerabilityMetadataInputSpec) (string, error) {
	funcName := "IngestVulnerabilityMetadata"
	err := helper.ValidateNoVul(vulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	err = helper.ValidateVulnerabilityIDInputSpec(vulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	return r.Backend.IngestVulnerabilityMetadata(ctx,
		model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityID)}, vulnerabilityMetadata)
}

// IngestBulkVulnerabilityMetadata is the resolver for the ingestBulkVulnerabilityMetadata field.
func (r *mutationResolver) IngestBulkVulnerabilityMetadata(ctx context.Context, vulnerabilities []*model.VulnerabilityInputSpec, vulnerabilityMetadataList []*model.VulnerabilityMetadataInputSpec) ([]string, error) {
	funcName := "IngestVulnerabilityMetadatas"
	if len(vulnerabilities) != len(vulnerabilityMetadataList) {
		return []string{}, gqlerror.Errorf("%v :: uneven vulnerabilities and vulnerabilityMetadata for ingestion", funcName)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	var lowercaseVulnInputList []*model.VulnerabilityInputSpec
	for _, v := range vulnerabilities {

		err := helper.ValidateNoVul(*v)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = helper.ValidateVulnerabilityIDInputSpec(*v)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(v.Type),
			VulnerabilityID: strings.ToLower(v.VulnerabilityID),
		}
		lowercaseVulnInputList = append(lowercaseVulnInputList, &lowercaseVulnInput)
	}
	return r.Backend.IngestBulkVulnerabilityMetadata(ctx, lowercaseVulnInputList, vulnerabilityMetadataList)
}

// VulnerabilityMetadata is the resolver for the vulnerabilityMetadata field.
func (r *queryResolver) VulnerabilityMetadata(ctx context.Context, vulnerabilityMetadataSpec model.VulnerabilityMetadataSpec) ([]*model.VulnerabilityMetadata, error) {
	funcName := "IngestVulnerabilityMetadata"
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if vulnerabilityMetadataSpec.Comparator != nil && vulnerabilityMetadataSpec.ScoreValue == nil {
		return []*model.VulnerabilityMetadata{}, gqlerror.Errorf("%v :: comparator cannot be set without a score value specified", funcName)
	}

	if vulnerabilityMetadataSpec.Vulnerability != nil {

		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if vulnerabilityMetadataSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		err := helper.ValidateVulnerabilitySpec(*vulnerabilityMetadataSpec.Vulnerability)
		if err != nil {
			return []*model.VulnerabilityMetadata{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnFilter := model.VulnerabilitySpec{
			ID:              vulnerabilityMetadataSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          vulnerabilityMetadataSpec.Vulnerability.NoVuln,
		}

		lowercaseVulnerabilityMetadataSpec := model.VulnerabilityMetadataSpec{
			ID:            vulnerabilityMetadataSpec.ID,
			Vulnerability: &lowercaseVulnFilter,
			ScoreType:     vulnerabilityMetadataSpec.ScoreType,
			ScoreValue:    vulnerabilityMetadataSpec.ScoreValue,
			Comparator:    vulnerabilityMetadataSpec.Comparator,
			Timestamp:     vulnerabilityMetadataSpec.Timestamp.UTC(),
			Origin:        vulnerabilityMetadataSpec.Origin,
			Collector:     vulnerabilityMetadataSpec.Collector,
		}
		return r.Backend.VulnerabilityMetadata(ctx, &lowercaseVulnerabilityMetadataSpec)
	} else {
		return r.Backend.VulnerabilityMetadata(ctx, &vulnerabilityMetadataSpec)
	}
}
