<script lang="ts">
    import CluesRules from './clues.rules.svelte';
    import { page } from '$app/stores';
    import { Button, Input } from '@mpiorowski/svelte-init';
    import type { Clue, User } from 'src/types';

    const id = $page.params.id;
    export let conn: WebSocket;
    export let user: User;

    let movieClue: Clue = {
        userId: user.id,
        word: 'movieClue',
        type: 'movie',
        guessed: false,
    };
    let personClue: Clue = {
        userId: user.id,
        word: 'personClue',
        type: 'person',
        guessed: false,
    };
    let placeClue: Clue = {
        userId: user.id,
        word: 'placeClue',
        type: 'place',
        guessed: false,
    };

    const onSendClues = () => {
        const clues = [movieClue, personClue, placeClue];
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

<form on:submit|preventDefault={onSendClues} class="mb-8">
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
    <Button type="ghost">Send clues</Button>
</form>
