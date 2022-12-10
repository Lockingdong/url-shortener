package application

import (
	"UrlShortener/src/adapter/output_port"
	"UrlShortener/src/domain/repository"
	"context"
	"reflect"
	"testing"
)

func TestCreateShortUrlCodeUseCase_Execute(t *testing.T) {
	type fields struct {
		createShortUrlInfoRepository repository.ICreateShortUrlCodeRepository
	}
	type args struct {
		ctx context.Context
		cmd *CreateShortUrlCodeCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateShortUrlCodeResponse
		wantErr bool
	}{
		{
			name: "TestCreateShortUrlCode",
			fields: fields{
				createShortUrlInfoRepository: output_port.NewShortUrlInfoMockRepository(nil),
			},
			args: args{
				ctx: context.Background(),
				cmd: &CreateShortUrlCodeCommand{
					OGUrl: "https://dongstudio.medium.com/",
				},
			},
			want: &CreateShortUrlCodeResponse{
				ShortUrlCode: "test_code",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &CreateShortUrlCodeUseCase{
				repository: tt.fields.createShortUrlInfoRepository,
			}
			got, err := u.Execute(tt.args.ctx, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortUrlCodeUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if reflect.TypeOf(got.ShortUrlCode) != reflect.TypeOf(tt.want.ShortUrlCode) {
				t.Errorf("CreateShortUrlCodeUseCase.Execute() = %v, want %v", reflect.TypeOf(got.ShortUrlCode), reflect.TypeOf(tt.want.ShortUrlCode))
			}
		})
	}
}
