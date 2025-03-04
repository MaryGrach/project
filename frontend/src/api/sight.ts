import { get } from '@api/base';
import { WithResponse, Sights, ReviewContent, Sight, Category } from 'src/types/api';

export async function getSights() {
  return await get('sights') as WithResponse<Sights>;
}

export async function getSight(sightId: number) {
  return await get(`sight/${sightId}`) as WithResponse<{ comments: ReviewContent[], sight: Sight }>;
}

export async function filterSights(hash: { [key: string]: unknown }) {
  let queryString = '';

  for (const key in hash) {
    if (hash.hasOwnProperty(key)) {
      const value = hash[key];
      queryString += `${encodeURIComponent(key)}=${encodeURIComponent(value as string)}&`;
    }
  }

  queryString = queryString.slice(0, -1);
  const url = `sights/search?${queryString}`;

  return get(url);
}

export async function getAdressByCoords(latitude : number, longitude : number) {
  return fetch(`https://nominatim.openstreetmap.org/reverse?lat=${latitude}&lon=${longitude}&zoom=18&accept-language=ru&format=json`);
}
