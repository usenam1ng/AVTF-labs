#include <iostream>

// Абстрактный класс персонажа
class Character {
public:
    virtual void setName(const std::string& name) = 0;
    virtual void setElement(const std::string& element) = 0;
    virtual void setRole(const std::string& role) = 0;
    virtual void setAttributes(const std::string& attributes) = 0;
    virtual void displayInfo() = 0;
};

// Абстрактная фабрика персонажей
class CharacterFactory {
public:
    virtual Character* createMage() = 0;
    virtual Character* createWarrior() = 0;
};

// Конкретный класс персонажа - Маг
class Mage : public Character {
private:
    std::string name;
    std::string element;
    std::string role;
    std::string attributes;

public:
    void setName(const std::string& name) override {
        this->name = name;
    }
    
    void setElement(const std::string& element) override {
        this->element = element;
    }
    
    void setRole(const std::string& role) override {
        this->role = role;
    }
    
    void setAttributes(const std::string& attributes) override {
        this->attributes = attributes;
    }
    
    void displayInfo() override {
        std::cout << "Character Info:\n";
        std::cout << "Name: " << name << "\n";
        std::cout << "Element: " << element << "\n";
        std::cout << "Role: " << role << "\n";
        std::cout << "Attributes: " << attributes << "\n";
        std::cout << "----------------------" << attributes << "\n";
        std::cout << "" << "\n";
    }
};

// Конкретный класс персонажа - Воин
class Warrior : public Character {
private:
    std::string name;
    std::string element;
    std::string role;
    std::string attributes;

public:
    void setName(const std::string& name) override {
        this->name = name;
    }

    void setElement(const std::string& element) override {
        this->element = element;
    }

    void setRole(const std::string& role) override {
        this->role = role;
    }

    void setAttributes(const std::string& attributes) override {
        this->attributes = attributes;
    }

    void displayInfo() override {
        std::cout << "Character Info:\n";
        std::cout << "Name: " << name << "\n";
        std::cout << "Element: " << element << "\n";
        std::cout << "Role: " << role << "\n";
        std::cout << "Attributes: " << attributes << "\n";
        std::cout << "----------------------" << attributes << "\n";
        std::cout << "" << "\n";
    }
};

// Конкретная фабрика создания персонажей стихии Земли
class EarthCharFactory : public CharacterFactory {
public:
    Character* createMage() override {
        return new Mage();
    }

    Character* createWarrior() override {
        return new Warrior();
    }
};

// Конкретная фабрика создания персонажей стихии Воды
class WaterCharFactory : public CharacterFactory {
public:
    Character* createMage() override {
        return new Mage();
    }

    Character* createWarrior() override {
        return new Warrior();
    }
};

int main() {
    std::cout << "Какой стихии персонажа вы хотите создать? (1 - Земли, 2 - Воды)\n";
    int a;
    std::cin >> a;

    if (a == 1){
        CharacterFactory* earthFactory = new EarthCharFactory();
        std::cout << "Какого класса персонажа вы хотите создать? (1 - Маг, 2 - Воин)\n";
        int b;
        std::cin >> b;
        if (b == 1){
            Character* mage = earthFactory->createMage();
            std::cout << "Введите имя персонажа: \n";
            std::string nm;
            std::cin >> nm;
            std::cout << "Свойства персонажа: \n";
            std::string sv;
            std::cin >> sv;
            mage->setName(nm);
            mage->setElement("Earth");
            mage->setRole("Mage");
            mage->setAttributes(sv);
            std::cout << "";
            mage->displayInfo();
        } else if (b == 2){
            Character* warrior = earthFactory->createWarrior();
            std::cout << "Введите имя персонажа: \n";
            std::string nm;
            std::cin >> nm;
            std::cout << "Свойства персонажа: \n";
            std::string sv;
            std::cin >> sv;
            warrior->setName(nm);
            warrior->setElement("Earth");
            warrior->setRole("Warrior");
            warrior->setAttributes(sv);
            std::cout << "";
            warrior->displayInfo();
        }
    } else if (a == 2){
        CharacterFactory* waterFactory = new WaterCharFactory();
        std::cout << "Какого класса персонажа вы хотите создать? (1 - Маг, 2 - Воин)\n";
        int b;
        std::cin >> b;
        if (b == 1){
            Character* mage = waterFactory->createMage();
            std::cout << "Введите имя персонажа: \n";
            std::string nm;
            std::cin >> nm;
            std::cout << "Свойства персонажа: \n";
            std::string sv;
            std::cin >> sv;
            mage->setName(nm);
            mage->setElement("Water");
            mage->setRole("Mage");
            mage->setAttributes(sv);
            std::cout << "";
            mage->displayInfo();
        } else if (b == 2){
            Character* warrior = waterFactory->createWarrior();
            std::cout << "Введите имя персонажа: \n";
            std::string nm;
            std::cin >> nm;
            std::cout << "Свойства персонажа: \n";
            std::string sv;
            std::cin >> sv;
            warrior->setName(nm);
            warrior->setElement("Water");
            warrior->setRole("Warrior");
            warrior->setAttributes(sv);
            std::cout << "";
            warrior->displayInfo();
        }
    }
    std::cout << std::endl;

    return 0;
}