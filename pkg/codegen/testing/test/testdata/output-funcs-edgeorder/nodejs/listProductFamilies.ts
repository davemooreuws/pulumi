// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * The list of product families.
 * API Version: 2020-12-01-preview.
 */
export function listProductFamilies(args: ListProductFamiliesArgs, opts?: pulumi.InvokeOptions): Promise<ListProductFamiliesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("myedgeorder::listProductFamilies", {
        "customerSubscriptionDetails": args.customerSubscriptionDetails,
        "expand": args.expand,
        "filterableProperties": args.filterableProperties,
        "skipToken": args.skipToken,
    }, opts);
}

export interface ListProductFamiliesArgs {
    /**
     * Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
     */
    customerSubscriptionDetails?: inputs.CustomerSubscriptionDetails;
    /**
     * $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
     */
    expand?: string;
    /**
     * Dictionary of filterable properties on product family.
     */
    filterableProperties: {[key: string]: inputs.FilterableProperty[]};
    /**
     * $skipToken is supported on list of product families, which provides the next page in the list of product families.
     */
    skipToken?: string;
}

/**
 * The list of product families.
 */
export interface ListProductFamiliesResult {
    /**
     * Link for the next set of product families.
     */
    readonly nextLink?: string;
    /**
     * List of product families.
     */
    readonly value: outputs.ProductFamilyResponse[];
}

export function listProductFamiliesOutput(args: ListProductFamiliesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListProductFamiliesResult> {
    return pulumi.output(args).apply(a => listProductFamilies(a, opts))
}

export interface ListProductFamiliesOutputArgs {
    /**
     * Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
     */
    customerSubscriptionDetails?: pulumi.Input<inputs.CustomerSubscriptionDetailsArgs>;
    /**
     * $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
     */
    expand?: pulumi.Input<string>;
    /**
     * Dictionary of filterable properties on product family.
     */
    filterableProperties: pulumi.Input<{[key: string]: pulumi.Input<pulumi.Input<inputs.FilterablePropertyArgs>[]>}>;
    /**
     * $skipToken is supported on list of product families, which provides the next page in the list of product families.
     */
    skipToken?: pulumi.Input<string>;
}
