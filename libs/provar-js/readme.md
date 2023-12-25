# @provar/provar-js

A JavaScript library for collecting user feedback and store them on Provar.se. This can be used on the client side as well as server side.

## Installation

```bash
yarn add @provar/provar-js
```

## Usage

```ts
import { ProvarClient } from '@provar/provar-js';

const client = new ProvarClient({ apiKey: 'your-api-key' });
await client.sendText('My feedback');
```
