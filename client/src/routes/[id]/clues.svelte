<script lang="ts">
    import { page } from "$app/stores";
    import type { Clue, User } from "src/types";

    const id = $page.params.id;
    export let conn: WebSocket;
    export let user: User;

    let movieClue: Clue = {
        userId: user.id,
        word: "movieClue",
        type: "movie",
        guessed: false,
    };
    let personClue: Clue = {
        userId: user.id,
        word: "personClue",
        type: "person",
        guessed: false,
    };
    let placeClue: Clue = {
        userId: user.id,
        word: "placeClue",
        type: "place",
        guessed: false,
    };

    const onSendClues = () => {
        const clues = [movieClue, personClue, placeClue];
        conn.send(
            JSON.stringify({
                type: "send-clues",
                data: JSON.stringify(clues),
                room: id,
                userId: user.id,
            })
        );
    };
</script>

<form on:submit|preventDefault={onSendClues}>
    <button type="submit">Submit</button>
    <label>
        Movie
        <input type="text" bind:value={movieClue.word} />
    </label>
    <label>
        Person
        <input type="text" bind:value={personClue.word} />
    </label>
    <label>
        Place
        <input type="text" bind:value={placeClue.word} />
    </label>
</form>
