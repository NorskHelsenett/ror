import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';
import { HighlightResult } from 'highlight.js';
import { Task } from '../../../core/models/task';

@Component({
  selector: 'app-config-task-detail',
  templateUrl: './config-task-detail.component.html',
  styleUrls: ['./config-task-detail.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class ConfigTaskDetailComponent implements OnInit {
  @Input()
  task: Task;

  response!: HighlightResult;

  constructor() {}

  ngOnInit(): void {
    return;
  }

  onHighlight(e: HighlightResult) {
    this.response = {
      language: e.language,
      relevance: e.relevance,
      _emitter: e._emitter,
      illegal: e.illegal,
      secondBest: e.secondBest,
      value: '{...}',
    };
  }
}
