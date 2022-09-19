// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Nursery extends pulumi.CustomResource {
    /**
     * Get an existing Nursery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Nursery {
        return new Nursery(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'plant:tree/v1:Nursery';

    /**
     * Returns true if the given object is an instance of Nursery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nursery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nursery.__pulumiType;
    }


    /**
     * Create a Nursery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NurseryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.varieties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'varieties'");
            }
            resourceInputs["sizes"] = args ? args.sizes : undefined;
            resourceInputs["varieties"] = args ? args.varieties : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Nursery.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Nursery resource.
 */
export interface NurseryArgs {
    /**
     * The sizes of trees available
     */
    sizes?: pulumi.Input<{[key: string]: pulumi.Input<enums.tree.v1.TreeSize>}>;
    /**
     * The varieties available
     */
    varieties: pulumi.Input<pulumi.Input<enums.tree.v1.RubberTreeVariety>[]>;
}
