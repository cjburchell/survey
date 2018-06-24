import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Observable} from "rxjs/index";

export interface Choice{
  id: string
  text: string
  value: number
}

export interface Question {
  id: string,
  text: string,
  type: string,
  choices: Choice[]
}

export interface Survey{
  name: string
  id: string
  questions: Question[]
}

export interface Result{
  surveyId: string,
  questionId: string,
  answer: string,
  count: number
}

export interface Answer{
  questionId: string,
  answer: string
}

export interface Count{
  count: number
}

@Injectable({
  providedIn: 'root'
})
export class SurveyService {

  constructor(private http: HttpClient) { }

  getSurvey(surveyId :string): Observable<Survey>{
    return this.http.get<Survey>(`/survey/${surveyId}`);
  }

  getCount(surveyId :string): Observable<Count>{
    return this.http.get<Count>(`/survey/${surveyId}/count`);
  }

  getResults(surveyId :string): Observable<Result[]>{
    return this.http.get<Result[]>(`/survey/${surveyId}/results`);
  }

  getGetResultsForQuestion(surveyId :string, questionId): Observable<Result[]>{
    return this.http.get<Result[]>(`/survey/${surveyId}/results/${questionId}`);
  }

  setAnswers(surveyId :string, answers :Answer[]){
    return this.http.post(`/survey/${surveyId}/answers`, answers);
  }
}
