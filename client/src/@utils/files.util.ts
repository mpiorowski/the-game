import { Client, gql, mutationStore, queryStore } from '@urql/svelte';
import type { CloudFile, FileType } from 'src/@types/files.type';
import { Config } from '../config';

export const insertFile = async (files: FileList, targetId: string, type: FileType) => {
  const formData = new FormData();
  for (const file of Array.from(files)) {
    formData.append('files', file);
  }
  const response = await fetch(
    `${Config.VITE_API_URL}/api/files?targetId=${targetId}&type=${type}`,
    {
      method: 'POST',
      body: formData,
      credentials: 'include',
    },
  );
  if (!response.ok) {
    throw new Error((await response.json()) as string);
  }
  const data = (await response.json()) as CloudFile;
  return data;
};

export const selectAllFilesGql = (client: Client, variables: { targetId: string | null }) =>
  queryStore<{ files: CloudFile[] }>({
    client: client,
    query: gql`
      query ($targetId: ID) {
        files(targetId: $targetId) {
          id
          targetId
          filename
          type
          url
        }
      }
    `,
    variables: variables,
    context: {
      additionalTypenames: ['File'],
    },
  });

export const deleteFileGql = (
  client: Client,
  variables: { targetId: string; fileId: string; filename: string },
) =>
  mutationStore<{ deleteFile: CloudFile }>({
    client: client,
    query: gql`
      mutation ($targetId: ID!, $fileId: ID!, $filename: String!) {
        deleteFile(targetId: $targetId, fileId: $fileId, filename: $filename) {
          id
        }
      }
    `,
    variables: variables,
  });
