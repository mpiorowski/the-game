export enum Team {
    A,
    B
}

export enum Catergory {
    Person = "person",
    Place = "place",
    Movie = "movie"
}

export type User = {
    id: string;
    nickname: string;
    team: number;
    ready: boolean;
    step: number;
};

export type Room = {
    id: string;
    created: string;
    updated: string;
    deleted: string;
    name: string;
    password: string;
}

export type Clue = {
    userId: string;
    word: string;
    type: string;
    guessed: boolean;
}

export type Round = {
    game: number;
    team: number;
    clue: Clue;
    user: User;
    nextUser: User;
    time: number;
}

export type Score = {
    game: number;
    teamClues: [Clue[], Clue[]];
};
