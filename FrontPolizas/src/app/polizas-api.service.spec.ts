import { TestBed } from '@angular/core/testing';

import { PolizasApiService } from './polizas-api.service';

describe('PolizasApiService', () => {
  let service: PolizasApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PolizasApiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
