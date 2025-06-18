import { cn } from '@/lib/utils';
import type { ButtonHTMLAttributes, ReactNode } from 'react';

type Props = {
  className?: string;
  children: ReactNode;
} & ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({ className, children, ...props }: Props) {
  return (
    <button
      data-slot="button"
      className={cn(
        'flex h-10 flex-row items-center justify-center gap-2',
        className
      )}
      {...props}
    >
      {' '}
      {children}
    </button>
  );
}
