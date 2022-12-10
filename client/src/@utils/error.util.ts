import { toast, ToastType } from '@mpiorowski/svelte-init';
import { Config } from '../config';
import { t } from "svelte-i18n"

export type PathError = Record<string, string>;

export const handleError = (error: unknown) => {
    if (Config.VITE_NODE_ENV === 'development') {
        console.error(error);
    }
    if (error && typeof error === 'object' && "code" in error && 'error' in error) {
        const GrpcError = error as { error: string, code: string }
        t.subscribe((value) => {
            toast(value(GrpcError.error), ToastType.ERROR)
        });
    } else {
        t.subscribe((value) => {
            toast(value('errors.somethingWentWrong'), ToastType.ERROR);
        });
    }
};
