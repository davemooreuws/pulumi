// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export interface ContainerArgs {
    brightness?: pulumi.Input<enums.ContainerBrightness>;
    color?: pulumi.Input<enums.ContainerColor | string>;
    material?: pulumi.Input<string>;
    size: pulumi.Input<enums.ContainerSize>;
}
/**
 * containerArgsProvideDefaults sets the appropriate defaults for ContainerArgs
 */
export function containerArgsProvideDefaults(val: ContainerArgs): ContainerArgs {
    return {
        ...val,
        brightness: (val.brightness) ?? 1,
    };
}
