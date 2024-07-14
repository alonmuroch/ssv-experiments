import {store} from '@ssvnetwork/sdk-ts';

const SLASHING_DB_KEY = "highest_slot";

export interface ExampleSigningData {
    slot: number;
    data: Uint8Array;
}

export function isSlashable(data: ExampleSigningData): boolean {
    if (data.slot > 100) {
        return true;
    }
    if (store.get(SLASHING_DB_KEY) > data.slot) {
        return true;
    }
    if (data.length > 1024) {
        return true
    }
    return false
}

export function sign(data: ExampleSigningData): Error {
    if (isSlashable(data)) {
        return new Error("slashable")
    }

    store.set(SLASHING_DB_KEY, data.slot);
    return null;
}