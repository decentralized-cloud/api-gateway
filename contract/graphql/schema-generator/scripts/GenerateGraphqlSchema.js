import fs from 'fs';
import path from 'path';
import { printSchema } from 'graphql';
import { getRootSchema } from '../src';

fs.writeFileSync(path.resolve(__dirname, '../../schema/schema.graphql'), printSchema(getRootSchema()));
