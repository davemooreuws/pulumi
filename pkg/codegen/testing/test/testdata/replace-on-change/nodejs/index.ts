// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { CatArgs } from "./cat";
export type Cat = import("./cat").Cat;
export const Cat: typeof import("./cat").Cat = null as any;

export { DogArgs } from "./dog";
export type Dog = import("./dog").Dog;
export const Dog: typeof import("./dog").Dog = null as any;

export { GodArgs } from "./god";
export type God = import("./god").God;
export const God: typeof import("./god").God = null as any;

export { NoRecursiveArgs } from "./noRecursive";
export type NoRecursive = import("./noRecursive").NoRecursive;
export const NoRecursive: typeof import("./noRecursive").NoRecursive = null as any;

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;

export { ToyStoreArgs } from "./toyStore";
export type ToyStore = import("./toyStore").ToyStore;
export const ToyStore: typeof import("./toyStore").ToyStore = null as any;

utilities.lazyLoad(exports, ["Cat"], () => require("./cat"));
utilities.lazyLoad(exports, ["Dog"], () => require("./dog"));
utilities.lazyLoad(exports, ["God"], () => require("./god"));
utilities.lazyLoad(exports, ["NoRecursive"], () => require("./noRecursive"));
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));
utilities.lazyLoad(exports, ["ToyStore"], () => require("./toyStore"));

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "example::Cat":
                return new Cat(name, <any>undefined, { urn })
            case "example::Dog":
                return new Dog(name, <any>undefined, { urn })
            case "example::God":
                return new God(name, <any>undefined, { urn })
            case "example::NoRecursive":
                return new NoRecursive(name, <any>undefined, { urn })
            case "example::ToyStore":
                return new ToyStore(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("example", "", _module)
pulumi.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
