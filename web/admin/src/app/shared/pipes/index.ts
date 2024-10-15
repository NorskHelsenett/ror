import { HighlighterPipe } from './highlighter.pipe';
import { FormatBytesPipe } from './format-bytes.pipe';
import { HighlightDifferencePipe } from './highlight-difference.pipe';
import { TimePipe } from './time.pipe';

export * from './format-bytes.pipe';
export * from './highlighter.pipe';
export * from './highlight-difference.pipe';

export const sharedPipes = [FormatBytesPipe, HighlighterPipe, HighlightDifferencePipe, TimePipe];
