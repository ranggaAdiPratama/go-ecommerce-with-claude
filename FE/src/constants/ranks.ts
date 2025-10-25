import type { Shop } from '../types/shop.types';

export const RANK_COLORS: Record<Shop['rank'], string> = {
    bronze: 'bg-amber-700',
    silver: 'bg-gray-400',
    gold: 'bg-yellow-500',
    platinum: 'bg-cyan-400',
};

export const RANK_TEXT_COLORS: Record<Shop['rank'], string> = {
    bronze: 'text-amber-700',
    silver: 'text-gray-600',
    gold: 'text-yellow-600',
    platinum: 'text-cyan-600',
};