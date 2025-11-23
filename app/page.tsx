import { headers, cookies } from "next/headers";

// Global in-memory counter for request ID (reset on server restart for demo)
let requestCounter = 0;

export default async function Home() {
  // Increment counter per request
  const requestId = ++requestCounter;

  // Get dynamic server data
  const headersStore = await headers();
  const requestMethod = headersStore.get('x-forwarded-method') || 'Unknown';
  const userAgent = headersStore.get('user-agent') || 'Unknown';
  const host = headersStore.get('host') || 'Unknown';
  const timestamp = new Date().toISOString();

  return (
    <main className="min-h-screen flex flex-col items-center justify-center px-8 py-16 space-y-12">
      <div className="max-w-2xl w-full grid grid-cols-1 md:grid-cols-2 gap-12">
        <div className="space-y-6 text-center animate-fade-in">
          <div className="flex flex-col items-center space-y-3">
            <svg className="w-12 h-12 text-foreground opacity-75" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h1 className="text-4xl md:text-5xl font-bold tracking-tight">Server Timestamp</h1>
            <p className="text-lg opacity-80">{timestamp}</p>
          </div>
        </div>

        <div className="space-y-6 text-center animate-fade-in">
          <div className="flex flex-col items-center space-y-3">
            <svg className="w-12 h-12 text-foreground opacity-75" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14" />
            </svg>
            <h1 className="text-4xl md:text-5xl font-bold tracking-tight">Request Method</h1>
            <p className="text-lg opacity-80">{requestMethod}</p>
          </div>
        </div>

        <div className="space-y-6 text-center animate-fade-in">
          <div className="flex flex-col items-center space-y-3">
            <svg className="w-12 h-12 text-foreground opacity-75" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h1 className="text-4xl md:text-5xl font-bold tracking-tight">Headers</h1>
            <div className="space-y-1 text-lg opacity-80">
              <p>User-Agent: {userAgent.slice(0, 40)}...</p>
              <p>Host: {host}</p>
            </div>
          </div>
        </div>

        <div className="space-y-6 text-center animate-fade-in">
          <div className="flex flex-col items-center space-y-3">
            <svg className="w-12 h-12 text-foreground opacity-75" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <h1 className="text-4xl md:text-5xl font-bold tracking-tight">Request ID</h1>
            <p className="text-lg opacity-80">#{requestId}</p>
          </div>
        </div>
      </div>
    </main>
  );
}
