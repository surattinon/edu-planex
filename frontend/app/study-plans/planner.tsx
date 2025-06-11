"use client";

import { useState } from 'react';

export default function Planner() {
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [isSubmitting, setSubmitting] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setSubmitting(true);
    try {
      const res = await fetch('/api/plan', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, description }),
      });
      if (res.ok) {
        // you could show a toast here
        setOpen(false);
        setTitle('');
        setDescription('');
      } else {
        console.error('Failed to apply');
      }
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <>
      <button
        onClick={() => setOpen(true)}
        className="px-6 py-3 bg-gradient-to-b from-primary/20 to-card text-card-foreground rounded-md border hover:bg-gray-700 transition"
      >
        New Plan
      </button>

      {open && (
        <div className="z-50 fixed inset-0 bg-black/60 flex items-center justify-center p-4">
          <div className="bg-gradient-to-b from-primary/2 to-card text-card-foreground rounded-md border shadow-lg backdrop-blur-sm w-full max-w-md p-6">
            <h2 className="text-2xl font-semibold mb-4">Plan Your Course</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-sm font-medium mb-1">
                  Course Title
                </label>
                <input
                  type="text"
                  required
                  value={title}
                  onChange={e => setTitle(e.target.value)}
                  className="w-full border rounded-lg px-3 py-2 focus:outline-none focus:ring"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">
                  Description
                </label>
                <textarea
                  required
                  value={description}
                  onChange={e => setDescription(e.target.value)}
                  className="w-full border rounded-lg px-3 py-2 h-24 resize-none focus:outline-none focus:ring"
                />
              </div>
              <div className="flex justify-end gap-3">
                <button
                  type="button"
                  onClick={() => setOpen(false)}
                  className="px-4 py-2 rounded-sm  bg-gradient-to-b from-primary/10 to-card text-card-foreground border hover:bg-red-900 transition"
                  disabled={isSubmitting}
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 rounded-sm  bg-gradient-to-b from-green-400/30 to-green-900/30 text-card-foreground border hover:bg-green-600 transition"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? 'Applying…' : 'Apply'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
}
