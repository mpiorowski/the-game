<script lang="ts">
    import CluesRules from './clues.rules.svelte';
    import { page } from '$app/stores';
    import { Button, Input, toast, ToastType } from '@mpiorowski/svelte-init';
    import type { Clue, User } from 'src/types';
    import { z } from 'zod';

    const id = $page.params.id;
    export let conn: WebSocket;
    export let user: User;

    let movieClue: Clue = {
        userId: user.id,
        word: '',
        type: '',
        guessed: false,
    };
    let personClue: Clue = {
        userId: user.id,
        word: '',
        type: '',
        guessed: false,
    };
    let placeClue: Clue = {
        userId: user.id,
        word: '',
        type: '',
        guessed: false,
    };

    const onSendClues = () => {
        const clues = [movieClue, personClue, placeClue];

        const schema = z.array(
            z.object({
                userId: z.string().uuid(),
                word: z.string().min(1),
                type: z.enum(['movie', 'person', 'place']),
                guessed: z.boolean(),
            })
        );
        if (
            !schema.safeParse(clues).success ||
            !clues.every((clue) => clue.word.length > 0) ||
            clues.some((clue) => clue.word.split(' ').length > 5)
        ) {
            toast('Invalid clues', ToastType.ERROR);
            return;
        }

        conn.send(
            JSON.stringify({
                type: 'send-clues',
                data: JSON.stringify(clues),
                room: id,
                userId: user.id,
            })
        );
    };

    const error = 'No more then 5 words.';
    let movieError = '';
    $: if (movieClue.word.split(' ').length > 5) {
        movieError = error;
    } else {
        movieError = '';
    }
    let personError = '';
    $: if (personClue.word.split(' ').length > 5) {
        personError = error;
    } else {
        personError = '';
    }
    let placeError = '';
    $: if (placeClue.word.split(' ').length > 5) {
        placeError = error;
    } else {
        placeError = '';
    }
</script>

{#if user.step === 2}
    <form
        on:submit|preventDefault={onSendClues}
        id="send-clues"
        class="mb-8 px-6"
    >
        <CluesRules />
        <Input
            type="text"
            label={'Movie / TV Show / Book / etc. (' +
                movieClue.word.split(' ').length +
                '/5)'}
            error={movieError}
            bind:value={movieClue.word}
        />
        <Input
            type="text"
            label={'Person (' + personClue.word.split(' ').length + '/5)'}
            bind:value={personClue.word}
            error={personError}
        />
        <Input
            type="text"
            label={'Place (' + placeClue.word.split(' ').length + '/5)'}
            bind:value={placeClue.word}
            error={placeError}
        />
        <div class="mt-4">
            <Button type="primary" form="send-clues">Send clues</Button>
        </div>
    </form>
{/if}
{#if user.step === 3}
    <h2 class="text-center">Waiting for other players to submit clues...</h2>
{/if}
