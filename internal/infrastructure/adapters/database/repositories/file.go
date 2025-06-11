package repositories

import (
	"encoding/json"
	"github.com/miloalej-dev/supermarket/internal/application/ports/outbound"
	"github.com/miloalej-dev/supermarket/internal/domain"
	"io"
	"os"
	"sync"
)

var fileRepoMutex sync.Mutex

type FileRepository struct {
	path string
}

func NewFileRepository() outbound.ProductRepository {
	return &FileRepository{
		path: "products_copy.json",
	}
}

func (f *FileRepository) readProducts() ([]domain.Product, error) {
	fileRepoMutex.Lock()
	defer fileRepoMutex.Unlock()
	file, err := os.OpenFile(f.path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []domain.Product{}, nil
	}
	var products []domain.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (f *FileRepository) writeProducts(products []domain.Product) error {
	fileRepoMutex.Lock()
	defer fileRepoMutex.Unlock()
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.path, data, 0666)
}

func (f *FileRepository) Save(product domain.Product) (domain.Product, error) {
	products, err := f.readProducts()
	if err != nil {
		return domain.Product{}, err
	}
	maxId := 0
	for _, p := range products {
		if p.Id > maxId {
			maxId = p.Id
		}
	}
	product.Id = maxId + 1
	products = append(products, product)
	if err := f.writeProducts(products); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (f *FileRepository) FindById(id int) *domain.Product {
	products, err := f.readProducts()
	if err != nil {
		return nil
	}
	for _, p := range products {
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func (f *FileRepository) FindAll() []domain.Product {
	products, err := f.readProducts()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

func (f *FileRepository) Update(id int, product domain.Product) *domain.Product {
	products, err := f.readProducts()
	if err != nil {
		return nil
	}
	for i, p := range products {
		if p.Id == id {
			product.Id = id
			products[i] = product
			if err := f.writeProducts(products); err != nil {
				return nil
			}
			return &products[i]
		}
	}
	return nil
}

func (f *FileRepository) UpdatePartial(id int, updates map[string]interface{}) *domain.Product {
	products, err := f.readProducts()
	if err != nil {
		return nil
	}
	for i, p := range products {
		if p.Id == id {
			// Apply updates
			if v, ok := updates["name"]; ok {
				p.Name = v.(string)
			}
			if v, ok := updates["quantity"]; ok {
				p.Quantity = int(v.(float64))
			}
			if v, ok := updates["code_value"]; ok {
				p.CodeValue = v.(string)
			}
			if v, ok := updates["is_published"]; ok {
				p.IsPublished = v.(bool)
			}
			if v, ok := updates["expiration"]; ok {
				p.Expiration = v.(string)
			}
			if v, ok := updates["price"]; ok {
				p.Price = v.(float64)
			}
			products[i] = p
			if err := f.writeProducts(products); err != nil {
				return nil
			}
			return &products[i]
		}
	}
	return nil
}

func (f *FileRepository) Delete(product domain.Product) error {
	products, err := f.readProducts()
	if err != nil {
		return err
	}
	idx := -1
	for i, p := range products {
		if p.Id == product.Id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return os.ErrNotExist
	}
	products = append(products[:idx], products[idx+1:]...)
	return f.writeProducts(products)
}

func (f *FileRepository) FindByPriceGreaterThan(price float64) []domain.Product {
	products, err := f.readProducts()
	if err != nil {
		return []domain.Product{}
	}
	var result []domain.Product
	for _, p := range products {
		if p.Price > price {
			result = append(result, p)
		}
	}
	return result
}
