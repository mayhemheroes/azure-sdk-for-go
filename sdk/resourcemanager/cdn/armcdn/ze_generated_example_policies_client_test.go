//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafListPolicies.json
func ExamplePoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyGet.json
func ExamplePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"MicrosoftCdnWafPolicy",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"MicrosoftCdnWafPolicy",
		armcdn.WebApplicationFirewallPolicy{
			Location: to.Ptr("WestUs"),
			Properties: &armcdn.WebApplicationFirewallPolicyProperties{
				CustomRules: &armcdn.CustomRuleList{
					Rules: []*armcdn.CustomRule{
						{
							Name:         to.Ptr("CustomRule1"),
							Action:       to.Ptr(armcdn.ActionTypeBlock),
							EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("CH")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRemoteAddr),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorGeoMatch),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.Ptr("windows")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRequestHeader),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorContains),
									Selector:        to.Ptr("UserAgent"),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.Ptr("<?php"),
										to.Ptr("?>")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableQueryString),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorContains),
									Selector:        to.Ptr("search"),
									Transforms: []*armcdn.TransformType{
										to.Ptr(armcdn.TransformTypeURLDecode),
										to.Ptr(armcdn.TransformTypeLowercase)},
								}},
							Priority: to.Ptr[int32](2),
						}},
				},
				ManagedRules: &armcdn.ManagedRuleSetList{
					ManagedRuleSets: []*armcdn.ManagedRuleSet{
						{
							RuleGroupOverrides: []*armcdn.ManagedRuleGroupOverride{
								{
									RuleGroupName: to.Ptr("Group1"),
									Rules: []*armcdn.ManagedRuleOverride{
										{
											Action:       to.Ptr(armcdn.ActionTypeRedirect),
											EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateEnabled),
											RuleID:       to.Ptr("GROUP1-0001"),
										},
										{
											EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateDisabled),
											RuleID:       to.Ptr("GROUP1-0002"),
										}},
								}},
							RuleSetType:    to.Ptr("DefaultRuleSet"),
							RuleSetVersion: to.Ptr("preview-1.0"),
						}},
				},
				PolicySettings: &armcdn.PolicySettings{
					DefaultCustomBlockResponseBody:       to.Ptr("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
					DefaultCustomBlockResponseStatusCode: to.Ptr(armcdn.PolicySettingsDefaultCustomBlockResponseStatusCode(200)),
					DefaultRedirectURL:                   to.Ptr("http://www.bing.com"),
				},
				RateLimitRules: &armcdn.RateLimitRuleList{
					Rules: []*armcdn.RateLimitRule{
						{
							Name:         to.Ptr("RateLimitRule1"),
							Action:       to.Ptr(armcdn.ActionTypeBlock),
							EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("192.168.1.0/24"),
										to.Ptr("10.0.0.0/24")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRemoteAddr),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorIPMatch),
									Transforms:      []*armcdn.TransformType{},
								}},
							Priority:                   to.Ptr[int32](1),
							RateLimitDurationInMinutes: to.Ptr[int32](0),
							RateLimitThreshold:         to.Ptr[int32](1000),
						}},
				},
			},
			SKU: &armcdn.SKU{
				Name: to.Ptr(armcdn.SKUNameStandardMicrosoft),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPatchPolicy.json
func ExamplePoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"MicrosoftCdnWafPolicy",
		armcdn.WebApplicationFirewallPolicyPatchParameters{
			Tags: map[string]*string{
				"foo": to.Ptr("bar"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyDelete.json
func ExamplePoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"rg1",
		"Policy1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
