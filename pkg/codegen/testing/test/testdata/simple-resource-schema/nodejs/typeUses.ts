// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

import {Resource} from "./index";

export class TypeUses extends pulumi.CustomResource {
    /**
     * Get an existing TypeUses resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TypeUses {
        return new TypeUses(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'example::TypeUses';

    /**
     * Returns true if the given object is an instance of TypeUses.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TypeUses {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TypeUses.__pulumiType;
    }

    public readonly bar!: pulumi.Output<outputs.SomeOtherObject | undefined>;
    public readonly baz!: pulumi.Output<outputs.ObjectWithNodeOptionalInputs | undefined>;
    public readonly foo!: pulumi.Output<outputs.Object | undefined>;

    /**
     * Create a TypeUses resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TypeUsesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["bar"] = args ? args.bar : undefined;
            resourceInputs["baz"] = args ? args.baz : undefined;
            resourceInputs["foo"] = args ? args.foo : undefined;
        } else {
            resourceInputs["bar"] = undefined /*out*/;
            resourceInputs["baz"] = undefined /*out*/;
            resourceInputs["foo"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TypeUses.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TypeUses resource.
 */
export interface TypeUsesArgs {
    bar?: pulumi.Input<inputs.SomeOtherObjectArgs>;
    baz?: pulumi.Input<inputs.ObjectWithNodeOptionalInputsArgs>;
    foo?: pulumi.Input<inputs.ObjectArgs>;
}
